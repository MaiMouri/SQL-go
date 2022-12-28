import { createSlice } from '@reduxjs/toolkit';
import { getAccounts } from "../api";


//createSliceを使ったら reducer を作成するだけで自動的に action type も定義してくれるし action creator も生成。
export const accountSlice = createSlice({
  //name -> action type の prefix
  name: "accounts",
  initialState: { loadingNow: false, error: null, accounts: [] },
  reducers: {
    //addUser -> action type
    // addAccount: (state, action) => {
    //   state.accounts = [action.payload, ...state.accounts]
    // },
    // updateAccount: (state, action) => {
    //   const { AccountId, Word, ImageUrl, Description } = action.payload
    //   const existingAccount = state.accounts.find((keyword) => keyword.AccountId === AccountId);
    //   console.log("upadtaAccount in slice" + existingAccount);
    //   console.log(AccountId, Word, ImageUrl, Description);
    //   // state.accounts = state.accounts.map((item) => item.AccountId === action.payload.AccountId);
    //   if (existingAccount) {
    //     existingAccount.Word = Word;
    //     existingAccount.Description = Description;
    //     existingAccount.ImageUrl = ImageUrl;
    //   }
    // },
    // deleteAccount: (state, action) => {
    //   state.accounts = state.accounts.filter((item) => item.AccountId !== action.payload.AccountId);
    // },
     // 通信を開始した時に呼ぶ関数
    fetchStart(state, action) {
      state.loadingNow = true;
      state.error = null;
    },
    // 通信が失敗した時に呼ぶ関数
    fetchFailure(state, action) {
      state.loading = false;
      state.error = action.payload;
    },
    // 通信が成功した時に呼ぶ関数
    fetchSuccess(state, action) {
      state.loadingNow = false;
      state.error = null;
      state.accounts = action.payload;
    },
  },
});

// Actions
export const { fetchStart, fetchFailure, fetchSuccess } = accountSlice.actions;
// export const { addAccount, updateAccount, deleteAccount, fetchStart, fetchFailure, fetchSuccess } = keywordSlice.actions;


// 外部からはこの関数を呼んでもらう
export const fetchItems = () => async dispatch => {
  try {
    dispatch(fetchStart());
    dispatch(fetchSuccess(await getAccounts()));
  } catch (error) {
    dispatch(fetchFailure(error.stack));
  }
};
// export const postAccount = (word) => async dispatch => {
//   try {
//     dispatch(addAccount(await createAccount(word)));
//   } catch (error) {
//     dispatch(fetchFailure(error.stack));
//   }
// };

// export const renewalAccount = (id) => async dispatch => {
//   try {
//     dispatch(updateAccount(await changeAccount(id)));
//   } catch (error) {
//     dispatch(fetchFailure(error.stack));
//   }
// };

// export const delAccount = (id) => async dispatch => {
//   try {
//     dispatch(deleteAccount(await removeAccount(id)));
//   } catch (error) {
//     dispatch(fetchFailure(error.stack));
//   }
// };

// Selectors keywords-storeのreducer nameと揃える
export const selectAccount = ({ accounts }) => accounts;

// Reducer(must be default export)
export default accountSlice.reducer;