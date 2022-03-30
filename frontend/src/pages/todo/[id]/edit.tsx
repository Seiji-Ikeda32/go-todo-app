import { NextPage, NextPageContext } from "next";
import { ChangeEvent, SetStateAction, useState } from "react";
import axios from "axios";

import Title from "../../../components/Title"
import Button from "../../../components/Button";

interface NullTime {
    Time: string;
    Valid: boolean;
}

interface Props {
    todo: {
        id: string;
        title: string;
        discription: string;
        is_completed: boolean;
        due_time: NullTime;
        created_at: NullTime;
        updated_at: NullTime;
    }
}

const TodoEdit: NextPage<Props> = ({todo}) => {
    const [title, setTitle] = useState(todo.title);
    const [discription, setDiscription] = useState(todo.discription);
    const [isCompleted, setIsCompleted] = useState(todo.is_completed);
    const [dueTimeStr, setDueTime] = useState(todo.due_time.Time);
    const [dueValid, setDueValid] = useState(todo.due_time.Valid);

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

    const UpdateTodo = () => {
        axios.put(`http://localhost:8080/todos/${todo.id}`, {
            title: title,
            discription: discription,
            is_completed: isCompleted,
            due_time: {Time: dueTimeStr, Valid: dueValid}
        })
        .then(res => {
            alert("todoを更新しました")
        })
    }

    return (
      <>
        <Title titleName="todo編集ページ" />
        
        <p>{todo.id}</p>
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

        <p>作成時間{todo.created_at.Time}({todo.created_at.Valid + ''})</p>
        <p>更新時間{todo.updated_at.Time}({todo.updated_at.Valid + ''})</p>

        <Button onClick={UpdateTodo}>todo更新</Button>

        <h2>todo削除</h2>
      </>
    )
}

TodoEdit.getInitialProps = async ({query: { id } }: NextPageContext) => {
    const res = await axios.get(`http://localhost:8080/todos/${id}`)
    const todo = await res.data
    return { todo }
}


export default TodoEdit;