
export default function Search({ handleSearchInput }) {

    return (
        <div className="flex justify-center">
            <div className="mb-3 xl:w-96">
                <label for="exampleSearch2" className="form-label inline-block mb-2 text-gray-700">Search</label>
                <input
                    onKeyUp={(e) => handleSearchInput(e.target.value)}
                    type="search"
                    className="
        form-control
        block
        w-full
        px-3
        py-1.5
        text-base
        font-normal
        text-gray-700
        bg-white bg-clip-padding
        border border-solid border-gray-300
        rounded
        transition
        ease-in-out
        m-0
        focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none
      "
                    placeholder="Type query"
                />
            </div>
        </div>
    )
}
