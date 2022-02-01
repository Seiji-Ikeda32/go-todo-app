import Link from "next/link"
import axios from "axios"
import { useState } from "react";

export default function HealthCheck() {

    const [message, setMesssage] = useState("");

    const getHealthCheck = () => {
        axios.get("http://localhost:8080/health")
        .then(res => {
            console.log(res.data)
            setMesssage(res.data.message)
        })
        .catch((error) => {
            "エラーが発生しました。"
        });
    }

    return (
      <>
        <h1>Health Checkページ</h1>

        <button onClick={getHealthCheck}>Health Check</button>
        <p>{message}</p>

        <h2>
          <Link href="/">
            <a>homeへ</a>
          </Link>
        </h2>
      </>
    )
}