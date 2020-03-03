import React, { useState } from "react"
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import Modal from 'react-bootstrap/Modal'
import './styles/Barris.css'
const config = require('../config.json')
const endpoint = config.serverEndpoint


const EditBarriModal = (props) => {
  const [barriName, setBarriName] = useState(null)
  const [url, setUrl] = useState(null)
    const editBarri = () => {
      const barriDomain = props.barriEdited.domain
      axios.post(endpoint+"/updateBarri?domain="+barriDomain+"&url="+url)
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
            {props.barriEdited.domain} - Editar
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group>
                <Form.Label>Nom</Form.Label>
                <Form.Control type="text" defaultValue={props.barriEdited.name} value={barriName} onChange={(e) => setBarriName(e.target.value)}/>
            </Form.Group>
            <Form.Group>
                <Form.Label>Url</Form.Label>
                <Form.Control type="text" defaultValue={props.barriEdited.url} value={url} onChange={(e) => setUrl(e.target.value)}/>
            </Form.Group>
            {/* <Form.Group>
              <Form.Label>Canal de Telegram: {props.barriEdited.telegramChannelId}</Form.Label>
            </Form.Group> */}
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