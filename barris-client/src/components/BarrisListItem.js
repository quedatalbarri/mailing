import React from "react"
import ListGroup from 'react-bootstrap/ListGroup'
import Button from 'react-bootstrap/Button'
import { FiEdit } from "react-icons/fi";
import './styles/Barris.css'

const BarrisListItem = (props) => {
    return(
        <ListGroup.Item>
              <div class="d-flex w-100 justify-content-between">
                {props.name}
                <div>
                    {!props.telegramToken && 
                        <Button variant="outline-primary" size="sm" className="qb-list-btn" onClick={() => {
                                props.setBarriEdited()
                                props.showTelegramModal()}
                            }
                        >
                        Conectar con Telegram
                    </Button>
                    }
                    <FiEdit color="grey"
                        onClick={() => {
                            props.setBarriEdited()
                            props.showEditBarriModal()}
                        }
                        className="qb-edit-icon"
                    />
                </div>
              </div>
              <div className="qb-list-url">{props.url}</div>
              {props.telegramToken && <div className="qb-list-url">Telegram Token: {props.telegramToken}</div>}
            </ListGroup.Item>
    )
 }

export default BarrisListItem