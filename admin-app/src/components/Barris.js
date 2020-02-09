/**Notice here that the useAuth0 hook is again being used, this time to retrieve the user's profile information 
 * (through the user property) and a loading property that can be used to display some kind of "loading" text 
 * or spinner graphic while the user's data is being retrieved.
 * In the UI for this component, the user's profile picture, name, and email address is being retrieved
 *  from the user property and displayed on the screen. */

import React, { Fragment } from "react"
import { useAuth0 } from "../react-auth0-spa"
import Form from 'react-bootstrap/Form'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Button from 'react-bootstrap/Button'
import Card from 'react-bootstrap/Card'

const Barris = () => {
  const { loading, user } = useAuth0();
  function handleSubmit() {
    alert('PostgreSQL or MongoDB?');
  }
  if (!user) {
    return <div>Registrat</div>;
  }
  else if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <Fragment>
      <h3>Barrios creados</h3>
      <p>Born</p>
      <Card>
        <Card.Body>
            <Card.Title>Crear barrio</Card.Title>
            <Form>
                <Form.Group as={Row} controlId="formPlaintextBarri">
                    <Form.Label column sm="2">
                    Barri
                    </Form.Label>
                    <Col sm="10">
                    <Form.Control defaultValue="born" />
                    </Col>
                </Form.Group>
                <Button variant="primary" type="submit" onClick={handleSubmit}>
                    Crear
                </Button>
            </Form>
            </Card.Body>
        </Card>
    </Fragment>
  );
};

export default Barris