"use client";
import { useRouter } from "next/navigation";
import { UserInfo } from "@/modules/auth_provider";
import React, { useContext, useEffect, useState } from "react";
import { AuthContext } from "@/modules/auth_provider";
import { setCookie } from 'cookies-next';

const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { authenticated } = useContext(AuthContext);
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  const router = useRouter();

  useEffect(() => {
    if (authenticated) {
      router.push("/");
      return;
    }
  }, [authenticated]);

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();
    try {
      const res = await fetch(`${apiUrl}/auth/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      const data = await res.json();
      console.log(data)
      if (res.ok) {
        const user: UserInfo = {
          username: data.username,
          id: data.id,
        };

        // localStorage.setItem("user_info", JSON.stringify(user));
        setCookie('is_authorized', true);

        // router.push('/');
        return router.push("/");
      }
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className="flex items-center justify-center min-w full min-h-screen">
      <form className="flex flex-col md:w-1/5">
        <div className="text-3xl font-bold text-center">
          <span className="text-blue-500"> Welcome!</span>
        </div>
        <input
          type="text"
          placeholder="email"
          className="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="password"
          className="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button
          className="p-3 mt-6 rounded-md bg-blue-500 font-bold text-white"
          type="submit"
          onClick={submitHandler}
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default LoginPage;
