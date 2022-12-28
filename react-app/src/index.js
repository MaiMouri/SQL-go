import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import './index.css';
import App from './App';
import ErrorPage from './components/ErrorPage';
import Home from './components/Home';
import Accounts from './components/Accounts';
import Tutorials from './components/Tutorials';

import { Provider } from 'react-redux';
import store from './redux/store/store'


const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {index: true, element: <Home /> },
      {
        path: "/accounts",
        element: <Accounts />,
      },
      {
        path: "/tutorials",
        element: <Tutorials />,
      },
    ]
  }
])


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Provider store={store}>
    <RouterProvider router={router} />
    </Provider>
  </React.StrictMode>
);

