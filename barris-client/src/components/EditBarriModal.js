import React from "react"
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Modal from 'react-bootstrap/Modal'
import './styles/Barris.css'
const config = require('../config.json')
const endpoint = config.serverEndpoint


const EditBarriModal = (props) => {

    const editBarri = () => {
      const barriName = props.barriEdited.name
      const url = document.getElementById('modalEditUrl').value
      const token = document.getElementById('modalEditTelegramToken').value
      axios.post(endpoint+"/updateBarri?name="+barriName+"&url="+url+"&telegramToken="+token)
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
            {props.barriEdited.name} - Editar
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group>
              <Form.Label>Url</Form.Label>
              <Form.Control type="text" defaultValue={props.barriEdited.url} id="modalEditUrl"/>
            </Form.Group>
            <Form.Group>
              <Form.Label>Token</Form.Label>
              <Form.Control type="text" defaultValue={props.barriEdited.telegramToken} id="modalEditTelegramToken"/>
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={props.onHide}>Cerrar</Button>
          <Button onClick={editBarri}>Guardar Cambios</Button>
        </Modal.Footer>
      </Modal>
    )
}

export default EditBarriModal