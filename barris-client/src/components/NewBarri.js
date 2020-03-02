/**Notice here that the useAuth0 hook is again being used, this time to retrieve the user's profile information 
 * (through the user property) and a loading property that can be used to display some kind of "loading" text 
 * or spinner graphic while the user's data is being retrieved.
 * In the UI for this component, the user's profile picture, name, and email address is being retrieved
 *  from the user property and displayed on the screen. */

import React, { Fragment, useState } from "react"
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Card from 'react-bootstrap/Card'
//const config = require('../config.json')
const endpoint = "http://localhost:1323"

const NewBarri = (props) => {

  const email = props.user.email
  const [barri, setBarri] = useState(null)
  const [url, setUrl] = useState(null)
  function handleSubmit(e) {
    e.preventDefault()
    if(!barri || !url) {
        return alert("campo Barri o url vacíos")
    }
    //axios.post(endpoint+"/addBarri?name="+barri+"&url="+url+"&telegramChannelId="+telegramChannelId+"&email="+email)
    axios.post(endpoint+"/barris?name="+barri+"&url="+url+"&email="+email)
        .then(res => {
            console.log(res.data.name);
            alert('Barri '+res.data.name+ " recibido en el server")
        })
        .catch(error => {
            alert(error.message)
        }) 
  }

  return (
    <Fragment>
      <Card className="qb-card">
        <Card.Body>
            <Card.Title>Crear barrio</Card.Title>
            <Form>
                <Form.Group>
                    <Form.Label>Barri*</Form.Label>
                    <Form.Control placeholder="ej: born" value={barri} onChange={(e) => setBarri(e.target.value)}/>
                </Form.Group>
                <Form.Group>
                    <Form.Label>Url*</Form.Label>
                    <Form.Control placeholder="barri_url" value={url} onChange={(e) => setUrl(e.target.value)}/>
                </Form.Group>
                {/* <Form.Group>
                    <Form.Label>Canal de Telegram</Form.Label>
                    <Form.Control placeholder="Identificador del canal de telegram" id="telegramChannelId"/>
                    <Form.Text className="text-muted">Opcional. Puedes añadirlo luego</Form.Text>
                </Form.Group> */}
                <Button variant="primary" type="submit" onClick={handleSubmit}>
                    Crear
                </Button>
            </Form>
            </Card.Body>
        </Card>
    </Fragment>
  )
}

export default NewBarri