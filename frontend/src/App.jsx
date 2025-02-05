import React, { useState, useEffect } from 'react'
import axios from 'axios'

function App() {
    const [tasks, setTasks] = useState([])
    const [title, setTitle] = useState('')

    useEffect(() => {
        axios.get('http://localhost:8080/tasks').then((response) => {
            setTasks(response.data)
        })
    }, [])

    const addTask = () => {
        axios
            .post('http://localhost:8080/tasks', { title, completed: false })
            .then((response) => {
                setTasks([...tasks, response.data])
                setTitle('')
            })
    }

    return (
        <div>
            <h1>Todo List</h1>
            <input value={title} onChange={(e) => setTitle(e.target.value)} />
            <button onClick={addTask}>Add Task</button>
            <ul>
                {tasks.map((task) => (
                    <li key={task.id}>{task.title}</li>
                ))}
            </ul>
        </div>
    )
}

export default App
