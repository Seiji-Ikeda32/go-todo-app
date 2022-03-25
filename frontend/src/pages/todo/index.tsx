import axios from "axios";
import Link from "next/link";
import { type } from "os";
import { useState } from "react";
import Router from "next/router";

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
      <h1>todo list</h1>

      <button onClick={getTodos}>getTodos</button>
      <ul>
        {todos.map((todo:Todo) =>
            <li key={todo.id}>
              <Link href={`/todo/${todo.id}`}>
                <p>{todo.id}</p>
              </Link>
              <p>タイトル：{todo.title}</p>
              <p>説明：{todo.discription}</p>
              <p>完了：{todo.is_completed + ''}</p>
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