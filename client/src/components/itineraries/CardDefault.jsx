function CardDefault() {
  return (
    <div className="flex items-center gap-3 pl-2">
      <div className="absolute -left-1.5 h-2.5 w-2.5 rounded-full bg-blue"></div>
      <div className="flex flex-grow gap-3 p-3 rounded-lg shadow-lg">
        <div className="w-16 h-16 rounded-lg bg-blue"></div>
        <div className="flex flex-col gap-1">
          <h5 className="text-sm font-bold">Lyf Funan</h5>
          <div className="flex gap-2">
            <div className="inline-block w-4 h-4 bg-yellow"></div>
            4.5
          </div>
          <div className="flex gap-2">
            <div className="inline-block w-4 h-4 bg-yellow"></div>
            Singapore
          </div>
        </div>
      </div>
    </div>
  );
}

export default CardDefault;
