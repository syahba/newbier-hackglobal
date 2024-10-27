import Main from '../layouts/Main'
import Logo from '/image/Logo.png'
import { useNavigate } from "react-router-dom";

function Index() {
  const navigate = useNavigate();
  const setSession = (id) => {
    localStorage.setItem("user", id)
    navigate("/home");
  }

  return (
    <Main>
      <div className="flex h-full w-full flex-col items-center justify-center gap-3">
        <div className="h-36 w-36">
          <img src={Logo} alt="logo"/>
        </div>
        <button className="rounded bg-blue text-white px-3 py-2 shadow-md" onClick={() => setSession(1)}>User 1</button>
        <button className="rounded bg-blue text-white px-3 py-2 shadow-md" onClick={() => setSession(2)}>User 2</button>
      </div>
    </Main>
  )
}

export default Index