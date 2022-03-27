import axios from "axios";
import Link from "next/link";
import { useState } from "react";
import Title from "../../components/Title"
import Button from "../../components/Button"

type Todo ={
  id: string;
  title: string;
  discription: string;
  is_completed: boolean;
  due_time: {Time: string; Valid:boolean;};
  created_at: {Time: string; Valid:boolean;};
  updated_at: {Time: string; Valid:boolean;};
}

const todo = () => {
  const [todos, setTodos] = useState([]);

  const getTodos = () => {
    axios.get("http://localhost:8080/todos/")
    .then((res) => {
      console.log(res.data)
      setTodos(res.data)
    })
    .catch((error) => {
      "エラーが発生しました。"
    });
  }

  return (
    <>
      <Title titleName="todo 一覧" />

      <Link href={`/todo/create`}>
          <h2>todo作成</h2>
        </Link>

      <Button buttonContent="todo一覧取得" onClick={getTodos} />
      <ul>
        {todos.map((todo:Todo) =>
            <li key={todo.id}>
              <Link href={`/todo/${todo.id}`}>
                <p>{todo.id}</p>
              </Link>
              <p>タイトル:{todo.title}</p>
              <p>説明:{todo.discription}</p>
              <p>完了:{todo.is_completed + ''}</p>
            </li>
        )}
      </ul>

      <h2>
        <Link href="/">
          <a>homeへ</a>
        </Link>
      </h2>
    </>
  )
}

export default todo;