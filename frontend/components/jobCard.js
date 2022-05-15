
export default function Search({ job }) {

    return (
        <a href={job.URL} target="_blank" rel="noreferrer" className="block p-6 max-w-sm bg-white rounded-lg border border-gray-200 shadow-md hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
            <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{job.title}</h5>
            <p className="font-bold text-gray-700 dark:text-gray-400">{job.companyName}</p>
            <p className="mb-2 font-bold text-gray-700 dark:text-gray-400">{job.source}</p>
            <p className="mb-2 font-bold text-gray-700 dark:text-gray-400">{job.location}</p>
            <p className="font-normal text-gray-700 dark:text-gray-400">{job.description}</p>
        </a>
    )
}
