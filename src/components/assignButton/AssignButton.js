import React from 'react'
import { Button } from 'react-bootstrap'
import { Link } from 'react-router-dom'

export const AssignButton = () => {
    return (
        <div className='mx-5 mt-5 d-flex flex-row-reverse'>
            <Link to="/admin/assign" className='d-flex flex-row-reverse text-white text-decoration-none'>
                <Button >
                    Assign Now
                </Button>
            </Link>
        </div>
    )
}
