import {React, useState} from 'react'
import { EditModal} from './inputs/Modals';

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCheckSquare,faCheck, faEdit} from '@fortawesome/free-solid-svg-icons'
library.add(faCheckSquare, faCheck)

const element = <FontAwesomeIcon icon={faCheck} />
const Task = (props) => {
    const [isEditerOpen, setEditerOpen] = useState(false);
    const {item, deleteTask, updateTodo} = props
    return (
        <div className="task">
        <dev className="task_title">
          
          <h4>{element}&nbsp;{item.title}</h4>
          <p style={{fontSize: "15px"}}>{item.description}</p>
        </dev>

        <div className='taskbuttons'>
            <a href="/" className="edit" onClick={(e)=> {
                  e.stopPropagation();
                  e.preventDefault();
                  setEditerOpen(true)
                }}>
                <FontAwesomeIcon icon={faEdit} color="#8A2BE2"/>
            </a>
            <a href="/" className="delete" onClick={(e)=> {
                  e.stopPropagation();
                  e.preventDefault();
                  deleteTask(item);
                }}>
                <FontAwesomeIcon icon={faCheckSquare} color="#8A2BE2"/>
            </a>
        </div>
        <EditModal open={isEditerOpen} todo={item} UpdateTodo={updateTodo} onClose={()=> setEditerOpen(false)}/>
      </div>
    )
}

export default Task