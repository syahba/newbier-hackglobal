function CardAction() {
  return (
    <div className="flex items-center gap-3 pl-2">
      <div className="absolute -left-1.5 h-2.5 w-2.5 rounded-full bg-blue"></div>
      <div className="flex flex-grow gap-3 rounded-lg shadow-lg">
        <div className="my-3 ml-3 h-16 w-16 rounded-lg bg-blue"></div>
        <div className="my-3 flex flex-grow flex-col justify-between">
          <h5 className="text-sm font-bold">Lyf Funan</h5>
          <div className="flex gap-2">
            <div className="inline-block h-4 w-4 bg-yellow"></div>
            4.5
          </div>
          <div className="flex gap-2">
            <div className="inline-block h-4 w-4 bg-yellow"></div>
            Singapore
          </div>
        </div>
        <div className="flex h-auto w-10 flex-col rounded-e border border-blue text-blue">
          <button className="flex flex-1 items-center justify-center hover:bg-blue hover:text-white">
            Del.
          </button>
        </div>
      </div>
    </div>
  );
}

export default CardAction;
