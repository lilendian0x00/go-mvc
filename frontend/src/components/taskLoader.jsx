import React, { Component } from 'react'
import axios from 'axios';
// Data
//import data from '../data/tasks.json';
import { AddModal } from './inputs/Modals';
import Task from './Task'


class API {
  constructor(debug=false) {
    this.debug = debug;

    this.endpoints = {
      "updateList": "/api/v1/list",
      "addTodo": "/api/v1/add",
      "updateTodo": "/api/v1/update",
      "deleteTodo": "/api/v1/remove/"
    }
  }

  GetFullEndpoint(action) {
    let fullEndpoint = "";
    (this.debug)? fullEndpoint="http://127.0.0.1" : fullEndpoint="";
    fullEndpoint += this.endpoints[action]
    return fullEndpoint;
  }
}

// true if you want to use 
var endpoints = new API((process.env.NODE_ENV)? true: false)

class TaskBox extends Component {
  constructor(props) {
    super(props);

    this.state = {
      taskList: [],
      isEditerOpen: false,
      isTodoMakerOpen: false,
    }
  }
  handleCheckboxChange = event => {
    // change dd in stat to checked artibute of handler element
    this.setState({ dd: event.target.checked });
  }

  triggerDelete(ID) {
    let taskList = [...this.state.taskList]
    

    this.state.taskList.forEach((todo, index) => {
      if (todo.ID === ID) {
        taskList.splice(index, 1);
        this.setState({taskList: taskList})
      }
    });
 }

  updateList = async() => {
    const resp = await axios.get(endpoints.GetFullEndpoint("updateList"))
    if (resp.status === 200) {
      this.setState({taskList: resp.data})
    }
  }

  addTodo = async(newTodo) => {
    const resp = await axios.post(endpoints.GetFullEndpoint("addTodo"), newTodo)
    if (resp.status === 200) {
      this.updateList()
      return true
    } else {
      return false
    }
  }

  updateTodo = async(newTodo) => {
    const resp = await axios.put(endpoints.GetFullEndpoint("updateTodo"), newTodo)
    if (resp.status === 200) {
      this.updateList()
      return true
    } else {
      return false
    }
  }

  
  deleteTodo = async(todo) => {
    console.log(todo)
    const resp = await axios.delete(endpoints.GetFullEndpoint("deleteTodo")+todo.ID)
    if (resp.status === 200) {
      this.triggerDelete(todo.ID)
    }
  }


  componentDidMount = async () => {
    await this.updateList()
  }
  
  render() {
    return (
      <div className="main">
          <h3 id="title">Tasks</h3>
          <div className="grid-container">
            <button className='ref' onClick={(e)=>{this.updateList()}}>LOAD-LIST</button>
            <button className='ref' onClick={(e)=>{this.setState({isTodoMakerOpen: true})}}>+</button>
            <AddModal open={this.state.isTodoMakerOpen} AddTodo={this.addTodo} onClose={()=> this.setState({isTodoMakerOpen: false})}></AddModal>
          </div>
          <div className="tasks">{
            this.state.taskList.map((item, i)=> (
              <Task key={item.id} item={item} deleteTask={this.deleteTodo} updateTodo={this.updateTodo}></Task>
          ))}
        </div>
      </div>
    );
  }

}


export default TaskBox;
