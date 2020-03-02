import React from "react"
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Modal from 'react-bootstrap/Modal'
import './styles/Barris.css'
const config = require('../config.json')
const endpoint = config.serverEndpoint


const TelegramModal = (props) => {

    const addToken = () => {
      const barriName = props.barriName
      const token = document.getElementById('modalTelegramToken').value
      if(token === "") return alert("No puede dejar el campo token vacío")
      axios.post(endpoint+"/updateBarriToken?name="+barriName+"&telegramToken="+token)
      .then(res => {
          console.log(res.data.name);
          alert('Barri '+res.data.name+ " recibido en el server")
          props.onHide()
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
            {props.barriName} - Telegram Token
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <p>
            Añadir explicación de para que sirve y cómo obtenerlo
          </p>
          <Form>
            <Form.Group>
              <Form.Label>Telegram Token</Form.Label>
              <Form.Control type="text" placeholder="Token" id="modalTelegramToken"/>
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={props.onHide}>Cerrar</Button>
          <Button onClick={addToken}>Añadir Token</Button>
        </Modal.Footer>
      </Modal>
    )
}

export default TelegramModal