import React, {useContext, useEffect, useState} from 'react';
import {privateApi} from "../services/api";
import {Profile as ProfileModel} from "typescript-axios";
import AuthContext from "./authContext";
import authHeaders from "../services/authHeaders";
import {Button, Card, CardGroup, Col, Row} from "react-bootstrap";

function People() {
  const [profiles, setProfiles] = useState<ProfileModel[]>([]);
  const auth = useContext(AuthContext);

  useEffect(() => {
    (async () => {
      try {
        const response = await privateApi.listProfiles({headers: authHeaders(auth.user)});
        setProfiles(response.data)
      } catch (e) {
        console.log(e)
      }
    })();
  }, []);

  return (
    <div>
      <h2>People</h2>
      <Row xs={1} md={3} className="g-4">
        {profiles.map(
          profile =>
            <Col key={profile.id}>
              <Card>
                <Card.Body>
                  <Card.Title>{profile.firstName} {profile.lastName}</Card.Title>
                  <Card.Subtitle>{profile.email}</Card.Subtitle>
                  <Card.Text>
                    {profile.city}, {profile.age} y. o.<br/>
                    {profile.interests}
                  </Card.Text>
                  <Button href={`/profile/${profile.id}`} variant="success" size="sm">View</Button>
                </Card.Body>
              </Card>
            </Col>
        )}
      </Row>
    </div>

  );
}

export default People;
