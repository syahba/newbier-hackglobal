import Main from '../layouts/Main'
import Logo from '/image/Logo.png'

function Index() {
  return (
    <Main>
      <div className="flex h-screen w-full justify-center bg-cyan-700">
        <section className="h-full w-96 bg-white">
          <div className="flex h-full w-full flex-col items-center justify-center gap-3">
            <div className="h-36 w-36">
              <img src={Logo} alt="logo"/>
            </div>
            <button className="rounded bg-blue text-white px-3 py-2 shadow-md">User 1</button>
            <button className="rounded bg-blue text-white px-3 py-2 shadow-md">User 2</button>
          </div>
        </section>
      </div>
    </Main>
  )
}

export default Index