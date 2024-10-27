function ButtonAction({text, onClick}) {
  return (
    <div>
      <button className="self-center px-5 py-3 text-sm rounded-lg shadow-lg w-fit bg-yellow hover:bg-blue" onClick={() => onClick && onClick()}>{text}</button>
    </div>
  );
}

export default ButtonAction;