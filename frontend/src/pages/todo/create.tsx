import { useState } from "react";
import axios from "axios";

import Title from "../../components/Title"
import Button from "../../components/Button";

const TodoCreate = () => {
    const [title, setTitle] = useState('');
    const [discription, setDiscription] = useState('');
    const [is_completed, setIsCompleted] = useState(false);
    const [due_time, setDueTime] = useState();
    const [due_valid, setDueValid] = useState(false);

    const handleTitle = (e: any) => {
        setTitle(e.target.value)
    }
    const handleDiscription = (e: any) => {
        setDiscription(e.target.value)
    }
    const handleDueTime = (e: any) => {
        setDueTime(e.target.value)
    }
    const handleStatus = () => {
        setIsCompleted(prevState => !prevState)
    }
    const handleDueStatus = () => {
        setDueValid(prevState => !prevState)
    }

    const CreateTodo = () => {
        axios.post("http://localhost:8080/todos/", {
            title: title,
            discription: discription,
            is_completed: is_completed,
            due_time: {Time: due_time, Valid: due_valid}
        })
        .then(res => {
            alert("todoを作成しました")
        })
    }

    return (
      <>
        <Title titleName="todo作成ページ" />
        
        <p>タイトル</p>
        <input
          onChange={(e) => handleTitle(e)}
          type={'text'}
          value={title}
        />

        <p>説明</p>
        <input
          onChange={(e) => handleDiscription(e)}
          type={'text'}
          value={discription}
        />

        <p>完了</p>
        <Button onClick={handleStatus}>
          {is_completed + ''}
        </Button>

        <p>期限</p>
        <Button onClick={handleDueStatus}>
          {due_valid + ''}
        </Button>
        <input
          onChange={(e) => handleDueTime(e)}
          type={'text'}
          value={due_time}
        />

        <Button onClick={CreateTodo}>
          todo作成
        </Button>

        <h2>todo削除</h2>
      </>
    )
}


export default TodoCreate;