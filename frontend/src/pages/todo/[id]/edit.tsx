import { NextPage, NextPageContext } from "next";
import { useState } from "react";
import axios from "axios";

import Title from "../../../components/Title"
import Button from "../../../components/Button";

interface Props {
    todo: {
        id: string;
        title: string;
        discription: string;
        is_completed: boolean;
        due_time: {Time: string; Valid:boolean;};
        created_at: {Time: string; Valid:boolean;};
        updated_at: {Time: string; Valid:boolean;};
    }
}

const TodoEdit: NextPage<Props> = ({todo}) => {
    const [title, setTitle] = useState(todo.title);
    const [discription, setDiscription] = useState(todo.discription);
    const [is_completed, setIsCompleted] = useState(todo.is_completed);
    const [due_time, setDueTime] = useState(todo.due_time.Time);
    const [due_valid, setDueValid] = useState(todo.due_time.Valid);

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

    const UpdateTodo = () => {
        axios.put(`http://localhost:8080/todos/${todo.id}`, {
            title: title,
            discription: discription,
            is_completed: is_completed,
            due_time: {Time: due_time, Valid: due_valid}
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
        <Button
          buttonContent={is_completed + ''}
          onClick={handleStatus}
        />

        <p>期限</p>
        <Button
          buttonContent={due_valid + ''}
          onClick={handleDueStatus}
        />
        <input
          onChange={(e) => handleDueTime(e)}
          type={'text'}
          value={due_time}
        />

        <p>作成時間{todo.created_at.Time}({todo.created_at.Valid + ''})</p>
        <p>更新時間{todo.updated_at.Time}({todo.updated_at.Valid + ''})</p>

        <Button
          buttonContent="todo更新"
          onClick={UpdateTodo}
        />

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