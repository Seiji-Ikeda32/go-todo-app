import axios from "axios";
import Link from "next/link";
import { useState } from "react";
import Title from "../../components/Title"
import Button from "../../components/Button"

interface NullTime {
  Time: string;
  Valid: boolean;
}

type Todo ={
  id: string;
  title: string;
  discription: string;
  isCompleted: boolean;
  dueTime: NullTime;
  createdAt: NullTime;
  updatedAt: NullTime;
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

      <Button onClick={getTodos}>todo一覧取得</Button>
      <ul>
        {todos.map((todo:Todo) =>
            <li key={todo.id}>
              <Link href={`/todo/${todo.id}`}>
                <p>{todo.id}</p>
              </Link>
              <p>タイトル:{todo.title}</p>
              <p>説明:{todo.discription}</p>
              <p>完了:{todo.isCompleted + ''}</p>
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