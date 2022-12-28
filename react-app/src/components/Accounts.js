import { Link } from "react-router-dom";
import { useState, useEffect } from "react";
import { fetchItems, selectAccount } from '../redux/feature/accountSlice';
import { useSelector, useDispatch } from "react-redux";
import { useTable } from "react-table";
import { columns, data } from "./tableData";

// react-table
import { ColumnDef, flexRender, getCoreRowModel, useReactTable } from '@tanstack/react-table';

const Accounts = () => {
  const { loadingNow, error, accounts } = useSelector(selectAccount);
  const [ columns, setColumns ] = useState([])

  const dispatch = useDispatch();
  let newColumns = [];

    // START FETCHING
    useEffect(() => {
      dispatch(fetchItems());
      console.log(accounts);
      // getKeys(accounts)
      // const columnsNames = Object.keys(accounts.accounts[0])
      

      // const columns = [
      //   { Header: columnsNames[0], accessor: columnsNames[0] },
      //   { Header: columnsNames[1], accessor: columnsNames[1] },
      //   { Header: columnsNames[2], accessor: columnsNames[2] }
      // ];
      // console.log(newColumns);
      // setColumns(newColumns);
      // newColumns = [];
    }, [dispatch]);
    // FINISH FETCHING
  

    const getKeys = (array) => {
      const columnsNames = Object.keys(array[0])
      console.log(columnsNames);
      for(let i = 0;i<columnsNames.length;i++) {
        // console.log(columnsNames[i]);
        newColumns.push({ Header: columnsNames[i], accessor: columnsNames[i] })
      }
    }


    // react-table
    const {
      getTableProps,
      getTableBodyProps,
      headerGroups,
      rows,
      prepareRow
    } = useTable({
      columns,
      accounts
    });
    // react-table
  

  return (
    <>
      <div className="text-center">
        <h2>Accounts</h2>
        <hr />
        <div>
        <table {...getTableProps()}>
            <thead>
              {headerGroups.map((headerGroup) => (
                <tr {...headerGroup.getHeaderGroupProps()}>
                  {headerGroup.headers.map((column) => (
                    <th {...column.getHeaderProps()}>
                      {column.render("Header")}
                    </th>
                  ))}
                </tr>
              ))}
            </thead>

            <tbody {...getTableBodyProps()}>
              {rows.map((row, i) => {
                prepareRow(row);
                return (
                  <tr {...row.getRowProps()}>
                    {row.cells.map((cell) => {
                      return (
                        <td {...cell.getCellProps()}>
                          {cell.render("Cell")}
                        </td>
                      )
                    })}
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      </div>
    </>
  );
};

export default Accounts;
