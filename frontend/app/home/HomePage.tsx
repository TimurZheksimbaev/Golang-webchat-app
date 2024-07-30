"use client";

import React, { useEffect, useState, useContext } from "react";
import { v4 as uuidv4 } from "uuid";
import { AuthContext } from "@/modules/auth_provider";
import { WebsocketContext } from "@/modules/websocket_provider";
import { useRouter } from "next/navigation";


const HomePage = () => {
  const [rooms, setRooms] = useState<{ id: string; name: string }[]>([]);
  const [roomName, setRoomName] = useState("");
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;
  const wsUrl = process.env.NEXT_PUBLIC_WEBSOCKET_URL
  const {user} = useContext(AuthContext)
  const {setConn} = useContext(WebsocketContext)
  const router = useRouter()

  const getRooms = async () => {
    try {
      const res = await fetch(`${apiUrl}/ws/rooms`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      const data = await res.json();
      if (res.ok) {
        setRooms(data);
      }
    } catch (err) {
      console.log(err);
    }
  };

  const joinRoom = async (roomId: string) => {
    const ws = new WebSocket(`${wsUrl}/ws/joinRoom/${roomId}?userId=${user.id}`);
    if (ws.OPEN) {
      setConn(ws)
      router.push("/app")
      return 
    }
  }

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();
    try {
      setRoomName("");
      const res = await fetch(`${apiUrl}/ws/createRoom`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({ id: uuidv4(), name: roomName }),
      });
      if (res.ok) {
        getRooms();
      }
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getRooms();
  }, []);

  return (
    <>
      <div className="my-8 px-4 md:mx-32 w-full h-full">
        <div className="flex justify-center mt-3 p-5">
          <input
            type="text"
            placeholder="room name"
            className="p-2 rounded-md border border-grey focus:outline-none focus:border-blue-500"
            value={roomName}
            onChange={(e) => setRoomName(e.target.value)}
          />
          <button
            onClick={submitHandler}
            className="bg-blue-500 border text-white rounded-md p-2 md:ml-4 font-sans"
          >
            Create Room
          </button>
        </div>

        <div className="mt-6">
          <div className="font-bold">Available Rooms</div>
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4 mt-6">
            {rooms.map((room, index) => (
              <div
                key={index}
                className="border border-blue-500 p-4 flex items-center rounded-md w-full"
              >
                <div className="w-full">
                  <div className="text-sm">Room</div>
                  <div className="text-blue-500 font-bold text-lg">
                    {room.name}
                  </div>
                </div>
                <div className="">
                  <button onClick={() => joinRoom(room.id)} className="px-4 text-white bg-blue-500 rounded-md">
                    {" "}
                    Join
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default HomePage;
