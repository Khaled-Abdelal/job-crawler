import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Search from '../components/search'
import JobCard from '../components/jobCard'
import { useMutation } from "react-query";
import React, { useState } from 'react';
import TablePagination from '@mui/material/TablePagination';

const APIRequest = async ({ searchTerm, pageInput, jobsPerPage }) => {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/jobs?searchTerm=${searchTerm}&size=${jobsPerPage}&from=${pageInput * jobsPerPage}`)
  const resJson = await res.json()
  if (!res.ok) throw resJson
  return resJson
}
export default function Home() {
  const [searchTerm, setSearchTerm] = useState("")
  const [jobsPerPage, setJobsPerPage] = useState(25)
  const [page, setPage] = useState(0)

  const mutation = useMutation(({ searchTerm, pageInput }) => APIRequest({ searchTerm, pageInput, jobsPerPage }))
  const handleSearchInput = (input = searchTerm, page = 0) => {
    setSearchTerm(input)
    setPage(page)
    if (input.trim() === "") {
      mutation.reset()
      return
    }
    mutation.mutate({ searchTerm: input, pageInput: page })
  }
  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="A job crawler site" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <Search handleSearchInput={handleSearchInput} />
        {
          mutation.isError ? (
            <div className="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800" role="alert">
              {mutation.error.message}
            </div>
          ) : null
        }
        {mutation.isSuccess && mutation.data.jobs &&
          <>
            <div className="container mx-auto">
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {mutation.data.jobs.map(job => {
                  return <JobCard job={job} key={job.URL}/>
                })}
              </div>
            </div>
            <TablePagination onRowsPerPageChange={(e) => setJobsPerPage(e.target.value)} rowsPerPage={jobsPerPage} count={mutation.data.total || 0} page={page} onPageChange={(_, page) => handleSearchInput(searchTerm, page)} />
          </>
        }
        {mutation.isSuccess && !mutation.data.jobs &&
          <>
            <div className="p-4 mb-4 text-sm text-blue-700 bg-blue-100 rounded-lg dark:bg-blue-200 dark:text-blue-800" role="alert">
              No results Found!
            </div>
          </>
        }
      </main>
    </div>
  )
}
