import { useEffect, useState } from 'react'
import { Container, Title, Card, Button, Text } from '@mantine/core'

export default function Orders(){
  const [orders, setOrders] = useState([])

  useEffect(()=>{
    fetch('/api/orders')
      .then(r=>r.json())
      .then(setOrders)
      .catch(()=>setOrders([]))
  },[])

  async function place(){
    const res = await fetch('/api/orders', {method:'POST', headers:{'content-type':'application/json'}, body: JSON.stringify({item:'Pizza', address:'123 Demo St'})})
    const data = await res.json()
    setOrders(prev=>[data, ...prev])
  }

  return (
    <Container style={{paddingTop:24}}>
      <Title order={2}>Orders</Title>
      <Button onClick={place} style={{marginBottom:12}}>Place Demo Order</Button>
      {orders.map(o=> (
        <Card key={o.id} shadow="xs" style={{marginBottom:8}}>
          <Text><strong>{o.item}</strong> — {o.address}</Text>
          <Text color="dimmed">Status: {o.status}</Text>
        </Card>
      ))}
    </Container>
  )
}
