import { Outlet } from 'react-router-dom';

function Scroll() {

  return (
    <section className="w-full h-auto flex justify-center bg-cyan-700 text-text">
      <section className="w-96 h-auto bg-white text-xs">
        <Outlet />
      </section>
    </section>
  )
}

export default Scroll