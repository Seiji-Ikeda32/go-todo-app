import axios from "axios";
import Link from "next/link";
import { type } from "os";
import { useState } from "react";

type Todo ={
  id: number;
  title: string;
  discription: string;
  is_completed: boolean;
  due_time: {Time: string; Valid:boolean;};
  created_at: {Time: string; Valid:boolean;};
  updated_at: {Time: string; Valid:boolean;};
}

export default function todo() {
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
      <h1>todo list</h1>

      <button onClick={getTodos}>getTodos</button>
      <ul>
        {todos.map((todo:Todo) =>
          <li key={todo.id}>
            <p>{todo.id}</p>
            <p>{todo.title}</p>
            <p>{todo.discription}</p>
            <p>{todo.is_completed + ''}</p>
            <p>{todo.due_time.Time}({todo.due_time.Valid + ''})</p>
            <p>{todo.created_at.Time}({todo.created_at.Valid + ''})</p>
            <p>{todo.updated_at.Time}({todo.updated_at.Valid + ''})</p>
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