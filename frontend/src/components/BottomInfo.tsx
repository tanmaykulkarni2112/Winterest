import type { JSX } from "react";

const BottomInfo = ():JSX.Element =>{
    return(
        <>
            <div className="bg-gray-800 h-150 1-full flex justify-around p-6 mt-20 px-16 ">
                <div className="text-4xl text-white m-20">
                    <i>Wintrest</i>
                </div>
                <div className="flex flex-col text-xl text-white m-20 p-6 px-16">
                    <div className="hover:bg-gray-500 rounded-lg m-2 p-3">
                        Github repo
                    </div>
                    <div className="m-2 p-3">
                        Profiles
                    </div>
                </div>
            </div>
        </>
    )
}

export default BottomInfo;