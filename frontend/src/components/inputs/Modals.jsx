import {useState} from 'react'
import './edit.scss'
import ReactDOM from 'react-dom'

export function EditModal({open, todo, onClose, UpdateTodo}) {
  const [title, setTitle] = useState(todo.title);
  const [description, setDescription] = useState(todo.description); 
  if (!open) {return null}

    return ReactDOM.createPortal(
    <>
    <div id='edit_modal_overlay'>
        <div id='edit_modal'>
          <button style={{backgroundColor: 'inherit', cursor: 'pointer'}} onClick={onClose}>X</button>
          <div id='mainForm'>
            <div className='title_form'>
                <label>Title:</label>
                <textarea name="title" value={title} onChange={(evt)=>setTitle(evt.target.value)}></textarea>
            </div>
          
            <div className='description_form'>
                <label>Description:</label>
                <textarea name="description" rows={5} value={description} onChange={(evt)=>setDescription(evt.target.value)}></textarea>
            </div>

            <div id='btn-sub'>
              <button onClick={async()=>{
                let newTodo = {id: todo.ID, title: title, description: description, done: false};
                let res = await UpdateTodo(newTodo);
                if (res) onClose();
            }}>Update</button>
            </div>

          </div>
        </div>
          
    </div>
    </>, document.getElementById("portal")
    )
}

export function AddModal({open, onClose, AddTodo}) {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState(''); 
  if (!open) {return null}

    return ReactDOM.createPortal(
    <>
    <div id='edit_modal_overlay'>
        <div id='edit_modal'>
          <button style={{backgroundColor: 'inherit', cursor: 'pointer'}} onClick={onClose}>X</button>
          <div id='mainForm'>
            <div className='title_form'>
                <label>Title:</label>
                <textarea name="title" value={title} onChange={(evt)=>setTitle(evt.target.value)}></textarea>
            </div>
          
            <div className='description_form'>
                <label>Description:</label>
                <textarea name="description" rows={5} value={description} onChange={(evt)=>setDescription(evt.target.value)}></textarea>
            </div>

            <div id='btn-sub'>
              <button onClick={async()=>{
                let newTodo = {title: title, description: description, done: false};
                let res = await AddTodo(newTodo);
                if (res) onClose();
            }}>Add</button>
            </div>

          </div>
        </div>
          
    </div>
    </>, document.getElementById("portal")
    )
}
