import { useRouter } from "next/router";
import { useState } from "react";
import axios from "axios";

type Todo ={
    id: string;
    title: string;
    discription: string;
    is_completed: boolean;
    due_time: {Time: string; Valid:boolean;};
    created_at: {Time: string; Valid:boolean;};
    updated_at: {Time: string; Valid:boolean;};
}

export default function tododetail() {
    const [todo, setTodo] = useState();
    const router = useRouter();
    const todoId = router.query.id

    const getTodo = () => {
      axios.get(`http://localhost:8080/todos/${todoId}`)
      .then((res) => {
        console.log(res.data)
        setTodo(res.data)
        console.log(todo)
      })
      .catch((error) => {
        "エラーが発生しました。"
      });
    }


    return (
      <>
        <h1>todo詳細ページ</h1>

        <button onClick={getTodo}>getTodo</button>
        
        {/* <p>{todo.id}</p>
        <p>タイトル：{todo.title}</p>
        <p>説明：{todo.discription}</p>
        <p>完了：{todo.is_completed + ''}</p>
        <p>期限{todo.due_time.Time}({todo.due_time.Valid + ''})</p>
        <p>作成時間{todo.created_at.Time}({todo.created_at.Valid + ''})</p>
        <p>更新時間{todo.updated_at.Time}({todo.updated_at.Valid + ''})</p> */}

        <h2>todo削除</h2>
        <h2>todo更新</h2>
      </>
    )
  }