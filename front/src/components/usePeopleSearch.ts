import {useContext, useEffect, useState} from "react";
import {privateApi} from "../services/api";
import authHeaders from "../services/authHeaders";
import AuthContext from "./authContext";
import axios, {Canceler} from "axios";
import {Profile} from "typescript-axios";

function usePeopleSearch(filterFirstName: string, filterLastName: string, offset: number, limit: number) {
  const auth = useContext(AuthContext);
  const [profiles, setProfiles] = useState<Profile[]>([]);
  const [hasMore, setHasMore] = useState(false);

  useEffect(() => {
    setProfiles([]);
  }, [filterFirstName, filterLastName]);

  useEffect(() => {
    let cancel: Canceler;

    privateApi
      .listProfiles(limit, offset, filterFirstName, filterLastName, {
        cancelToken: new axios.CancelToken(c => cancel = c),
        headers: authHeaders(auth.user),
      })
      .then(r => {
        if (r.data.entities) {
          setProfiles(prevProfiles => {
            return [...prevProfiles, ...r.data.entities!];
          });
        }

        setHasMore(r.data.hasMore || false);
      }).catch(e => {
        if (axios.isCancel(e)) {
          return;
        }

        throw e;
      })
    ;

    return () => cancel();
  }, [filterFirstName, filterLastName, offset]);

  return {profiles, hasMore};
}

export default usePeopleSearch;
