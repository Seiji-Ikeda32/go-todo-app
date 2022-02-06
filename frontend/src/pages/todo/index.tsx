import Link from "next/link";
// import axios from "axios";
// import { useState } from "react";

export default function todoList() {

    return (
      <>
        <h1>todo list</h1>

        <h2>
          <Link href="/">
            <a>homeへ</a>
          </Link>
        </h2>
      </>
    )
}