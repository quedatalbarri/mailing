/**Notice here that the useAuth0 hook is again being used, this time to retrieve the user's profile information 
 * (through the user property) and a loading property that can be used to display some kind of "loading" text 
 * or spinner graphic while the user's data is being retrieved.
 * In the UI for this component, the user's profile picture, name, and email address is being retrieved
 *  from the user property and displayed on the screen. */

import React, { Fragment } from "react"
import { useAuth0 } from "../react-auth0-spa"
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Button from 'react-bootstrap/Button'
import Card from 'react-bootstrap/Card'
//const config = require('../config.json')
const endpoint = "http://localhost:1323"

const NewBarri = () => {
  function handleSubmit(e) {
    e.preventDefault()
    const barri = document.getElementById('barriName').value
    const url = document.getElementById('barriUrl').value
    axios.post(endpoint+"/addBarri?name="+barri+"&url="+url)
        .then(res => {
            console.log(res.data.name);
            alert('Barri '+res.data.name+ " recibido en el server")
        })
        .catch(error => {
            debugger
        }) 
  }

  return (
    <Fragment>
      <Card>
        <Card.Body>
            <Card.Title>Crear barrio</Card.Title>
            <Form>
                <Form.Group as={Row}>
                    <Form.Label column sm="2">
                    Barri
                    </Form.Label>
                    <Col sm="10">
                    <Form.Control defaultValue="born" id="barriName"/>
                    </Col>
                </Form.Group>
                <Form.Group as={Row}>
                    <Form.Label column sm="2">
                    Url
                    </Form.Label>
                    <Col sm="10">
                    <Form.Control defaultValue="barri_url" id="barriUrl"/>
                    </Col>
                </Form.Group>
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