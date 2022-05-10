import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Search from '../components/search'
import JobCard from '../components/jobCard'
import { useMutation } from "react-query";
import React, { useState } from 'react';
import TablePagination from '@mui/material/TablePagination';

export default function Home() {
  const [searchTerm, setSearchTerm] = useState("")
  const [jobsPerPage, setJobsPerPage] = useState(25)
  const [page, setPage] = useState(0)

  const mutation = useMutation(({searchTerm, pageInput}) =>
    fetch(`http://localhost:8081/api/jobs?searchTerm=${searchTerm}&size=${jobsPerPage}&from=${pageInput * jobsPerPage}`)
    .then(res => res.json())
  )
  const handleSearchInput = (input=searchTerm, page=0) => {
    if(input.trim === "") return
    setSearchTerm(input)
    setPage(page)
    mutation.mutate({searchTerm:input, pageInput:page})
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
        {mutation.isSuccess && mutation.data.jobs &&
          <>
            <div className="container mx-auto">
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {mutation.data.jobs.map(job => {
                  return <JobCard job={job}/>
                })}
              </div>
            </div>
            <TablePagination onRowsPerPageChange={(e)=> setJobsPerPage(e.target.value)} rowsPerPage={jobsPerPage} count={mutation.data.total || 0} page={page} onPageChange={(_, page)=> handleSearchInput(searchTerm, page)} />
          </>
        }
      </main>
    </div>
  )
}
