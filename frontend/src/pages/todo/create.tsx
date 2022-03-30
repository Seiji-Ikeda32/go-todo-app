import { ChangeEvent, useState } from "react";
import axios from "axios";

import Title from "../../components/Title"
import Button from "../../components/Button";

const TodoCreate = () => {
    const [title, setTitle] = useState('');
    const [discription, setDiscription] = useState('');
    const [isCompleted, setIsCompleted] = useState(false);
    const [dueTimeStr, setDueTime] = useState("0001-01-01T00:00:00Z");
    const [dueValid, setDueValid] = useState(false);

    const handleTitle = (e: ChangeEvent<HTMLInputElement>) => {
        setTitle(e.target.value)
    }
    const handleDiscription = (e: ChangeEvent<HTMLInputElement>) => {
        setDiscription(e.target.value)
    }
    const handleDueTime = (e: ChangeEvent<HTMLInputElement>) => {
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
            is_completed: isCompleted,
            due_time: {Time: dueTimeStr, Valid: dueValid}
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
          {isCompleted + ''}
        </Button>

        <p>期限</p>
        <Button onClick={handleDueStatus}>
          {dueValid + ''}
        </Button>
        <input
          onChange={(e) => handleDueTime(e)}
          type={'text'}
          value={dueTimeStr}
        />

        <Button onClick={CreateTodo}>
          todo作成
        </Button>

        <h2>todo削除</h2>
      </>
    )
}


export default TodoCreate;