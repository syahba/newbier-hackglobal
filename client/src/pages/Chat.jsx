import { useEffect, useState } from "react";
import Scroll from "../layouts/Scroll"

function ChatRoom(){
  const userId = localStorage.getItem("user")
  const [result, setResult] = useState([]);
  const [message, setMessage] = useState("");

  useEffect(() => {
    try {
      const fetchData = async () => {
        try {
            const response = await fetch(`http://localhost:8000/api/chat`);
            const data = await response.json();
            setResult(data)
        } catch (error) {
            console.log(error)
        }
      }

      fetchData()
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  }, []);

  const sendMessage = async () => {
    try {
      await fetch({method: "POST", url: "http://localhost:8000/api/chat", body: {
        created_by: userId,
        message: message
      }});
    } catch (error) {
        console.log(error)
    } finally {
      setMessage("")
    }
  }

  return (
    <Scroll>
      <div className="flex flex-col gap-4 px-6 pt-4 mt-2 mb-4">
        {
          result.map(elm => {
            if (elm.created_by == userId) {
              return (
                <div className="self-end px-3 py-2 text-white rounded w-fit max-w-80 bg-blue">
                  <p>Hello~ ^^</p>
                </div>
              )
            } else {
              return (
                <div className="self-start px-3 py-2 rounded w-fit max-w-80 bg-gray">
                  <p>Hi! I saw you wanted to go to Universal Studio tomorrow. to go to Universal Studio tomorrow. to go to Universal Studio tomorrow.</p>
                </div>
              )
            }
          })
        }
      </div>

      <div className="absolute bottom-0 h-auto px-6 pb-6 text-sm bg-white rounded w-96">
        <div className="flex items-center justify-between w-full h-auto shadow-md">
          <input type="text" placeholder="Enter your message here" className="flex-grow p-3" value={message} onChange={e => setMessage(e.target.value)} />
          <button className="w-auto px-4 py-2 text-sm rounded h-fit" onClick={() => sendMessage()}>Send</button>
        </div>
      </div>
    </Scroll>
  )
}

export default ChatRoom