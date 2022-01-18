import React, {useContext, useEffect, useState} from 'react';
import {privateApi} from "../services/api";
import {Profile as ProfileModel} from "typescript-axios";
import AuthContext from "./authContext";
import authHeaders from "../services/authHeaders";
import {Button, Card, Col, Row, Spinner} from "react-bootstrap";
import InfiniteScroll from "react-infinite-scroll-component";

function People() {
  const [profiles, setProfiles] = useState<ProfileModel[]>([]);
  const [lastProfileId, setLastProfileId] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const auth = useContext(AuthContext);

  useEffect(() => {
    (async () => {
      await fetchProfiles();
    })();
  }, []);

  const fetchProfiles = async () => {
    const response = await privateApi.listProfiles(12, lastProfileId, {headers: authHeaders(auth.user)});

    if (response.data.entities) {
      setProfiles(profiles.concat(response.data.entities));
      setLastProfileId(response.data.entities[response.data.entities.length - 1].id!);
    }

    if (!response.data.hasMore) {
      setHasMore(false);
    }
  };

  return (
    <div>
      <h2>People</h2>
      <InfiniteScroll
        dataLength={profiles.length}
        next={fetchProfiles}
        hasMore={hasMore}
        loader={<Spinner animation="border"/>}
      >
        <Row xs={1} md={2}>
          {profiles.map(
            profile =>
              <Col key={profile.id} className={"g-4"}>
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
      </InfiniteScroll>
    </div>
  );
}

export default People;
