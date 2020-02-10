package main

import (
	"time"
	"strconv"
	"bytes"
	"io/ioutil"
	"text/template"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"golang.org/x/net/context"
    "github.com/hanzoai/gochimp3"
	"log"
	"os"
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


func getChimp() *gochimp3.API {
	apiKey := os.Getenv("MAILCHIMP_API_KEY")
	client := gochimp3.New(apiKey)
	return client
}

type Event struct {
	Summary string
	Description string
}

type EmailContent struct {
	Metadata map[string]string
	Events []Event
}


func makeEmailContent(events *calendar.Events) *EmailContent {
	es := make([]Event, len(events.Items))

	for i, item := range events.Items {
		es[i] = Event{item.Summary, item.Description}
	}

	// TODO: Get dates and time!
	// date := item.Start.DateTime
	// if date == "" {
		// date = item.Start.Date
	// }

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


func sendEmail(client *gochimp3.API, listId string, segmentId int, templatePath string, content *EmailContent) (bool, error) {

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


func main(){
	cal := os.Getenv("ABEJA_CALENDAR")
	templatePath := os.Getenv("ABEJA_TEMPLATE")
	listId := os.Getenv("ABEJA_LIST_ID")
	segmentId, err := strconv.Atoi(os.Getenv("ABEJA_SEGMENT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	srv := getCalendarService()
	client := getChimp()

	events, err := getUpcomingEvents(srv, cal)
	content := makeEmailContent(events)

	sent, err := sendEmail(client, listId, segmentId, templatePath, content)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Abeja has sent the email: %v", sent)
}
