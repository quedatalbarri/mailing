/**Notice here that the useAuth0 hook is again being used, this time to retrieve the user's profile information 
 * (through the user property) and a loading property that can be used to display some kind of "loading" text 
 * or spinner graphic while the user's data is being retrieved.
 * In the UI for this component, the user's profile picture, name, and email address is being retrieved
 *  from the user property and displayed on the screen. */

import React, { Fragment, useState } from "react"
import { useAuth0 } from "../react-auth0-spa"
import axios from 'axios'
import NewBarri from './NewBarri'
import BarrisListItem from './BarrisListItem'
import EditBarriModal from './EditBarriModal'
import TelegramModal from './TelegramModal'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import ListGroup from 'react-bootstrap/ListGroup'

import './styles/Barris.css'
const config = require('../config.json')
const endpoint = config.serverEndpoint


const Barris = () => {
  const { loading, user } = useAuth0()
  const [barrisList, setBarrisList] = useState(null)
  const [loadingBarris, setLoadingBarris] = useState(false)
  const [userBarrisList, setUserBarrisList] = useState(null)
  const [loadingUserBarris, setLoadingUserBarris] = useState(false)
  const [telegramModalShow, setTelegramModalShow] = React.useState(false)
  const [editBarriModalShow, setEditBarriModalShow] = React.useState(false)
  const [barriEdited, setBarriEdited] = React.useState(false)

  const getUserBarris = () => {
    setLoadingUserBarris(true)
    axios.get(endpoint+"/getBarris?email="+user.email)
    .then(res => {
        console.log(res.data.barris)
        const listItems = res.data.barris.map((b) => {
          return <BarrisListItem 
                  name={b.name}
                  telegramChannel={b.telegramChannelId}
                  url={b.url}
                  setBarriEdited={() => setBarriEdited(b)}
                  showTelegramModal={() => setTelegramModalShow(true)}
                  showEditBarriModal={() => setEditBarriModalShow(true)}
                />
        })
        setUserBarrisList(<ListGroup variant="flush" className="qb-list">{listItems}</ListGroup>)
        setLoadingUserBarris(false)
    })
    .catch(error => {
        debugger
    })
  }

  const getBarris = () => {
    setLoadingBarris(true)
    axios.get(endpoint+"/barris")
    .then(res => {
        console.log(res.data.barris)
        const listItems = res.data.barris.map((b) => {
            return <ListGroup.Item>{b.name}<div className="qb-list-url">{b.url}</div></ListGroup.Item>
        })
        setBarrisList(<ListGroup variant="flush" className="qb-list">{listItems}</ListGroup>)
        setLoadingBarris(false)
    })
    .catch(error => {
        debugger
    })
  }

  if (!user) {
    return <Fragment>
        <Row className="mt-5">
        <div>Registra't</div>
        </Row>
      </Fragment>
  }
  else if (loading) {
    return <div>Loading...</div>;
  }
  else if (!barrisList && !loadingBarris) {
    getBarris()
  }
  else if (!userBarrisList && !loadingUserBarris) {
    getUserBarris()
  }
  return (
    <Fragment>
      <Row className="mt-5">
        <Col md>
          <Row>
            <Col>
              <h4>Els meus barris</h4>
              {userBarrisList}
            </Col>
          </Row>
          <Row className="mt-3">
            <Col>
              <h4>Tots els barris</h4>
              {barrisList}
            </Col>
          </Row>
        </Col>
        <Col md>
          <NewBarri user={user}/>
        </Col>
      </Row>
      <EditBarriModal
        show={editBarriModalShow}
        onHide={() => setEditBarriModalShow(false)}
        barriEdited = {barriEdited}
      />
      <TelegramModal
        show={telegramModalShow}
        onHide={() => setTelegramModalShow(false)}
        barriName = {barriEdited.name}
      />
    </Fragment>
  );
};

export default Barris