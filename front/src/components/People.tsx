import React, {useState} from 'react';
import {Button, Card, Col, Form, Row, Spinner} from "react-bootstrap";
import InfiniteScroll from "react-infinite-scroll-component";
import usePeopleSearch from "./usePeopleSearch";

function People() {
  const limit = 30;
  const [query, setQuery] = useState("");
  const [offset, setOffset] = useState(0);
  const {profiles, hasMore} = usePeopleSearch(query, offset, limit);

  const handleSearch = (e: any) => {
    setQuery(e.target.value);
    setOffset(0);
  };

  const handleNext = () => {
    setOffset(prevOffset => prevOffset + limit);
  }

  return (
    <div>
      <h2>People</h2>
      <Form.Control
        className={"mt-4 mb-4"}
        type="text"
        placeholder="Start typing first name or last name..."
        onChange={handleSearch}
      />
      <InfiniteScroll
        dataLength={profiles.length}
        next={handleNext}
        hasMore={hasMore}
        loader={<Spinner animation="border"/>}
        scrollThreshold={0.95}
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
