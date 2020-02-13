/**Notice here that the useAuth0 hook is again being used, this time to retrieve the user's profile information 
 * (through the user property) and a loading property that can be used to display some kind of "loading" text 
 * or spinner graphic while the user's data is being retrieved.
 * In the UI for this component, the user's profile picture, name, and email address is being retrieved
 *  from the user property and displayed on the screen. */

import React, { Fragment, useState } from "react"
import { useAuth0 } from "../react-auth0-spa"
import axios from 'axios'
import NewBarri from './NewBarri'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import ListGroup from 'react-bootstrap/ListGroup'
import './styles/Barris.css'
//const config = require('../config.json')
const endpoint = "http://localhost:1323"


const Barris = () => {
  const { loading, user } = useAuth0()
  const [barrisList, setBarrisList] = useState(null);
  const getBarris = () => {
    axios.get(endpoint+"/getBarris")
    .then(res => {
      debugger
        console.log(res.data.barris)
        const listItems = res.data.barris.map((b) => {
          return <ListGroup.Item>{b.name}<div className="qb-list-url">{b.url}</div></ListGroup.Item>
        })
        setBarrisList(
          <ListGroup variant="flush" className="qb-list">{listItems}</ListGroup>)
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
  else if (!barrisList) {
    getBarris()
  }
  return (
    <Fragment>
      <Row className="mt-5">
        <Col md>
          <h3>Barrios creados</h3>
          {barrisList}
        </Col>
        <Col md>
          <NewBarri user={user}/>
        </Col>
      </Row>
    </Fragment>
  );
};

export default Barris