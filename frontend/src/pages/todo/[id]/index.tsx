import axios from "axios";
import { NextPage, NextPageContext } from "next";
import Link from "next/link";
import { useRouter } from 'next/router';

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


const Tododetail: NextPage<Props> = ({todo}) => {
    const router = useRouter();

    const deleteTodo = () => {
      if(confirm("このtodoを削除しますか？") ) {
        axios.delete(`http://localhost:8080/todos/${todo.id}`)
        .then(res => {
            alert("todoを削除しました")
            router.push('/todo')
        })
      }
      else {
        alert("削除を中止しました")
      }
    }

    return (
      <>
        <Title titleName="todo詳細ページ" />
        
        <Link href={`/todo/${todo.id}/edit`}>
          <h2>todo編集</h2>
        </Link>

        <p>{todo.id}</p>
        <p>タイトル:{todo.title}</p>
        <p>説明:{todo.discription}</p>
        <p>完了:{todo.is_completed + ''}</p>
        <p>期限{todo.due_time.Time}({todo.due_time.Valid + ''})</p>
        <p>作成時間{todo.created_at.Time}({todo.created_at.Valid + ''})</p>
        <p>更新時間{todo.updated_at.Time}({todo.updated_at.Valid + ''})</p>

        <Button
          buttonContent={'todo削除'}
          onClick={deleteTodo}
        />
      </>
    )
}

Tododetail.getInitialProps = async ({query: { id } }: NextPageContext) => {
    const res = await axios.get(`http://localhost:8080/todos/${id}`)
    const todo = await res.data
    return { todo }
}


export default Tododetail;