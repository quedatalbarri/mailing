/**Notice here that the useAuth0 hook is again being used, this time to retrieve the user's profile information 
 * (through the user property) and a loading property that can be used to display some kind of "loading" text 
 * or spinner graphic while the user's data is being retrieved.
 * In the UI for this component, the user's profile picture, name, and email address is being retrieved
 *  from the user property and displayed on the screen. */

import React, { Fragment } from "react"
import { useAuth0 } from "../react-auth0-spa"
import NewBarri from './NewBarri'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import ListGroup from 'react-bootstrap/ListGroup'
import './styles/BarrisListItem.css'

const Barris = () => {
  const { loading, user } = useAuth0()

  if (!user) {
    return <div>Registrat</div>;
  }
  else if (loading) {
    return <div>Loading...</div>;
  }
  return (
    <Fragment>
      <Row className="mt-5">
        <Col>
          <h3>Barrios creados</h3>
          <p>Todo- get list from database. list aqui:</p>
          <ListGroup variant="flush" className="qb-list">
            <ListGroup.Item>Born<div className="qb-list-url">"https://barcelona.us4.list-manage.com/subscribe?u=aafafb6a3fe6cd8bb1c071405&id=be49b3c618"</div></ListGroup.Item>
            <ListGroup.Item>Gotic</ListGroup.Item>
            <ListGroup.Item>Dapibus ac facilisis in</ListGroup.Item>
          </ListGroup>
        </Col>
        <Col>
          <NewBarri/>
        </Col>
      </Row>
    </Fragment>
  );
};

export default Barris