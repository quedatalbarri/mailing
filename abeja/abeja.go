package main

import (
	"time"
	"bytes"
	"io/ioutil"
	"text/template"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"golang.org/x/net/context"
    "github.com/hanzoai/gochimp3"
	"github.com/caarlos0/env/v6"
	"log"
)

func getCalendarService() *calendar.Service {
	client, err := google.DefaultClient(context.Background(), calendar.CalendarReadonlyScope)
	if err != nil {
		log.Print(err)
	}

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	return srv
}

func getUpcomingEvents(srv *calendar.Service, cal string) (*calendar.Events, error) {
	t := time.Now().Format(time.RFC3339)

	events, err := srv.Events.
		List(cal).
			ShowDeleted(false).
				SingleEvents(true).
					TimeMin(t).
						MaxResults(10).
							OrderBy("startTime").
								Do()

	if err != nil {
		return events, err
	}
	return events, nil
}


func getChimp(key string) *gochimp3.API {
	client := gochimp3.New(key)
	return client
}

type Event struct {
	Summary string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	Location string `json:"location,omitempty"`
	HtmlLink string `json:"link,omitempty"`
	DateTime string `json:"datetime,omitempty"`

	// TODO: Add link to association that created the event???
	// created by in the calendar... ???
}



type Week struct {
	// que week??
	// que barri???
	Events []Event `json:"events,omitempty"`
}

type EmailContent struct {
	Metadata map[string]string
	Events []Event
}

func MakeEvents(events *calendar.Events) []Event {
	es := make([]Event, len(events.Items))

	for i, item := range events.Items {
		es[i] = Event{
			item.Summary,
			item.Description,
			item.Location,
			item.HtmlLink,
			item.Start.DateTime,
		}
	}

	return es
}


func makeEmailContent(events *calendar.Events) *EmailContent {
	es := MakeEvents(events)

	metadata := map[string]string{"foo": "barrr"}
	return &EmailContent{metadata, es}
}


func getTemplate(path string, content *EmailContent) string {
	tpl, _ := ioutil.ReadFile(path)

	var out bytes.Buffer
	t, _ := template.New("email").Parse(string(tpl))
	t.Execute(&out, content)
	return out.String()
}


func updateCampaign(client *gochimp3.API, campaign string, templatePath string, content *EmailContent) error {
	html := getTemplate(templatePath, content)
	req := gochimp3.CampaignContentUpdateRequest{}
	req.Html = html

	_, err := client.UpdateCampaignContent(campaign, &req)
	return err
}

func createCampaign(client *gochimp3.API, listId string, segmentId int) (*gochimp3.CampaignResponse, error) {

	segment := gochimp3.CampaignCreationSegmentOptions{segmentId, gochimp3.CONDITION_MATCH_ALL, []string{}}
	recipients := gochimp3.CampaignCreationRecipients{listId, segment}
	settings := gochimp3.CampaignCreationSettings{}
	settings.SubjectLine = "Aquesta Setmana al Born"
	settings.FromName = "Queda't al Barri"
	settings.ReplyTo = "quedatalbarri@gmail.com"

	req := gochimp3.CampaignCreationRequest{gochimp3.CAMPAIGN_TYPE_REGULAR, recipients, settings, gochimp3.CampaignTracking{}}
	return client.CreateCampaign(&req)
}


func emailer(client *gochimp3.API, listId string, segmentId int, templatePath string, content *EmailContent) (bool, error) {

	campaign, err := createCampaign(client, listId, segmentId)

	if err != nil {
		log.Fatal(err)
	}
	c := campaign.ID

	err = updateCampaign(client, c, templatePath, content)
	if err != nil {
		return false, err
	}
	return client.SendCampaign(c, &gochimp3.SendCampaignRequest{c})
}


type Config struct {
	Calendar string `env:"ABEJA_CALENDAR,required"`
	TemplatePath string `env:"ABEJA_TEMPLATE,required"`
	ListID string `env:"ABEJA_LIST_ID,required"`
	SegmentID int `env:"ABEJA_SEGMENT_ID,required"`
	MailchimpKey string `env:"MAILCHIMP_API_KEY,required"`
}

func getConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return cfg
}

func sendEmail(cnf Config, events *calendar.Events) (bool, error){
	client := getChimp(cnf.MailchimpKey)
	content := makeEmailContent(events)
	return emailer(client, cnf.ListID, cnf.SegmentID, cnf.TemplatePath, content)
}

func main(){
	cnf := getConfig()
	srv := getCalendarService()
	events, err := getUpcomingEvents(srv, cnf.Calendar)
	sent, err := sendEmail(cnf, events)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Abeja has sent the email: %v", sent)
}
