import React, { Fragment, useState } from "react"
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Modal from 'react-bootstrap/Modal'
import './styles/Barris.css'
import './styles/TelegramModal.css'
const config = require('../config.json')
const endpoint = config.serverEndpoint


const TelegramModal = (props) => {

    const [channel, setChannel] = useState(null)
    const saveBarriChannel = (channel) => {
      const barriName = props.barriName
      if(channel === "") return alert("No puede dejar el campo canal vacío")
      axios.post(endpoint+"/updateBarriChannel?name="+barriName+"&TelegramChannelId="+channel)
      .then(res => {
          console.log(res.data.name);
          alert('Barri '+res.data.name+ " recibido en el server")
          props.onHide()
      })
      .catch(error => {
          alert(error.message)
      }) 
    }
    const validateBotInChannel = (e) => {
      e.preventDefault()
      if(channel === "" || !channel) return alert("Has dejado el campo canal vacío")
      axios.get(endpoint+"/getChatMember/"+channel)
      .then(res => {
          console.log(res.data)
          if (res.data.ok) {
            if (res.data.Result.Status === "administrator") {
              alert("Perfecto")
              saveBarriChannel(channel)   //TODO
              props.onHide()
            }
            else {
              alert("El bot no ha sido añadido como administrador a este canal")
            }
          }
          else {
            alert("Este canal no existe.")
          }
          
      })
      .catch(error => {
          alert(error.message)
      }) 
    }

    return (
        <Modal
        {...props}
        size="lg"
        aria-labelledby="contained-modal-title-vcenter"
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title id="contained-modal-title-vcenter">
            {props.barriName} - Telegram
          </Modal.Title>
        </Modal.Header>
          <Fragment>
            <Modal.Body>
              <p>Para conectar tu barrio con telegram, debes crear un canal en Telegram dónde se enviaran los eventos.</p>
              <Row className="qb-steps">
                <Col>
                  <h3>1<hr/></h3>
                  Crea un canal público desde tu app de telegram (Por ejemplo: quedatal{props.barriName})
                </Col>
                <Col>
                  <h3>2<hr/></h3>
                  <Form.Group>
                    <Form.Label>Indicanos el nombre del canal que has creado.</Form.Label>
                    <Form.Control type="text" placeholder={"ej: quedatal"+props.barriName} value={channel} onChange={(e) => setChannel(e.target.value)}/>
                  </Form.Group>
                </Col>
                <Col>
                  <h3>3<hr/></h3>
                  Añade, como administrador de tu canal, al bot <b>@PruebaQuedatBot</b>
                </Col>
              </Row>
              <Form onSubmit={validateBotInChannel}>
                {/* <Button onClick={props.onHide}>Cerrar</Button>{' '} */}
                <Button className="float-right" type="submit">Validar</Button>
              </Form>
            </Modal.Body> 
        </Fragment>
      </Modal>
    )
}

export default TelegramModal