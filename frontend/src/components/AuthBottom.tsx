import type { JSX } from "react"

type elements = {
    text: string,
    linktext: string,
    onClick: () => void
}

const AuthBottom = ({ text, linktext, onClick }: elements): JSX.Element => {
    return (
        <div className="text-gray-300 font-medium text-center flex">
            <div>
                {text}
            </div>
            <div onClick={onClick} className=" text-gray-500 hover:text-black">
                {linktext}
            </div>
        </div>
    )
}

export default AuthBottom;