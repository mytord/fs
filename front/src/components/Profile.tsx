import React, {useContext, useEffect, useState} from 'react';
import {privateApi as api} from "../services/api";
import {Profile as ProfileModel} from "typescript-axios";
import AuthContext from "./authContext";
import authHeaders from "../services/authHeaders";
import {Card} from "react-bootstrap";
import {useParams} from "react-router-dom";

function Profile() {
  const [profile, setProfile] = useState<ProfileModel | undefined>(undefined);
  const auth = useContext(AuthContext);
  const { id } = useParams();

  useEffect(() => {
    (async () => {
      try {
        let response: any;

        if (id) {
          response = await api.getProfile(Number(id), {headers: authHeaders(auth.user)});
        } else {
          response = await api.getCurrentProfile({headers: authHeaders(auth.user)});
        }

        setProfile(response.data)
      } catch (e) {
        console.log(e)
      }
    })();
  }, []);

  return (
    <div>
      <h2>{id ? "Profile" : "Your Profile"}</h2>
      <Card>
        {profile && <Card.Body>
          <Card.Title>{profile.firstName} {profile.lastName}</Card.Title>
          <Card.Subtitle>{profile.email}</Card.Subtitle>
          <Card.Text>
            {profile.city}, {profile.age} y. o.<br/>
            {profile.interests}
          </Card.Text>
        </Card.Body>}
      </Card>
    </div>
  );
}

export default Profile;
