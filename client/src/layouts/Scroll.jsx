
function Scroll(props) {

  return (
    <section className="flex justify-center w-full h-auto min-h-screen bg-cyan-700 text-text">
      <section className="h-auto text-xs bg-white w-96">
      {props.children}
      </section>
    </section>
  )
}

export default Scroll