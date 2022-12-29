import logo from './logo.svg';
import './App.css';
import { useForm } from 'react-hook-form';
import { useEffect } from "react";
import { Link, Outlet, useNavigate } from "react-router-dom";
import { fetchItems, selectAccount } from './redux/feature/accountSlice';
import { useSelector, useDispatch } from "react-redux";

function App() {
  const { register, handleSubmit, formState: { errors } } = useForm();
  // const { register, handleSubmit } = useForm();
  const onSubmit = (data) => {
    console.log(data)
    
  };

  console.log(errors);

  const navigate = useNavigate();
  //useSelectorでstoreの中のstateにアクセスできる。usersはreducer名
  const { loadingNow, error, accounts } = useSelector(selectAccount);
  const dispatch = useDispatch();

  // START FETCHING
  useEffect(() => {
    dispatch(fetchItems());
  }, [dispatch]);
  // FINISH FETCHING

  return (
    <div className="App">
      <header>SQL-GO</header>
      <div className="container">
      <div className="row">
        <div className="col">
          <h1 className="mt-3">Learn SQL 🖌️</h1>
        </div>
        <div className="col text-end">

        </div>
        <hr className="mb-3"></hr>
      </div>

      <div className="row">
          <div className="col-md-2">
            <nav>
              <div className='list-group'>
                <Link to="/" className="list-group-item list-group-item-action">
                  Home
                </Link>
                <Link to="/accounts" className="list-group-item list-group-item-action">
                  Accounts
                </Link>
                <Link to="/tutorials" className="list-group-item list-group-item-action">
                  Tutorials
                </Link>
                <Link to="/" className="list-group-item list-group-item-action">
                  Home
                </Link>
              </div>
            </nav>
          </div>
        
          <div className="col-md-10">
            <Outlet context={{dispatch, fetchItems}}/>
              {/* <Alert message={alertMessage} className={alertClassName} />
              
              <Outlet
              context={{
                jwtToken,
                setJwtToken,
                setAlertClassName,
                setAlertMessage,
                toggleRefresh,
              }}
            /> */}
            <form onSubmit={handleSubmit(onSubmit)}>
              <textarea {...register("query", {required: true})} className="form-control"/>

              <input type="submit" className="btn btn-primary" value="Send"/>
            </form>
          </div>
        </div>
    </div>
    </div>
    
  );
}

export default App;
