import React, { Fragment, useState } from "react"
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Modal from 'react-bootstrap/Modal'
import './styles/Barris.css'
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
      if(channel === "" || !channel) return alert("No puede dejar el campo canal vacío")
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
              <h5>Creación del canal en telegram donde se enviaran los eventos</h5>
              <p>Para ello, debes:
                <div>- Crear un canal público. (Lo puedes llamar quedatal{props.barriName}) </div>
                <div>- Añadir, como administrador de tu canal, el bot <b>@PruebaQuedatBot</b> </div>
              </p>
              <Form onSubmit={validateBotInChannel}>
                <Form.Group>
                  <Form.Label>Canal</Form.Label>
                  <Form.Control type="text" placeholder="ej. quedatalborne" value={channel} onChange={(e) => setChannel(e.target.value)}/>
                </Form.Group>
                <Button onClick={props.onHide}>Cerrar</Button>{' '}
                <Button type="submit">Validar</Button>
              </Form>
            </Modal.Body> 
        </Fragment>
      </Modal>
    )
}

export default TelegramModal