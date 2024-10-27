

function Main(props) {

  return (
    <section className="w-full h-screen flex justify-center bg-cyan-700 text-text">
      <section className="w-96 h-full bg-white text-xs">
        {props.children}
      </section>
    </section>
  )
}

export default Main