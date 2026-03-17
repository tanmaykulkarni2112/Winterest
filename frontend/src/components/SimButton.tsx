type ButtonProps = {
  text: string
  colour?: string
  textColour?: string
  onClick?: ()=>void
}

const Button = ({ text, colour, onClick, textColour }: ButtonProps)=> {
  return (
    <div>
      <button
        type="button"
        onClick={onClick}
        className={`${textColour} ${colour} box-border border border-transparent hover:bg-danger-strong focus:ring-4 focus:ring-danger-medium shadow-xs font-medium leading-5 rounded-base text-sm px-4 py-2.5 focus:outline-none
        rounded-xl p-6`}
      >
        {text}
      </button>
    </div>
  )
}

export default Button