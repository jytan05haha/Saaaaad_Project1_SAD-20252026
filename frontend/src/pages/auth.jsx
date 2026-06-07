import { Button, Container, TextInput, Title } from '@mantine/core'
import { useState } from 'react'

export default function Auth(){
  const [name, setName] = useState('')

  function login(){
    localStorage.setItem('token','demo');
    location.reload();
  }

  return (
    <Container size="xs" style={{paddingTop:40}}>
      <Title order={2}>Food Delivery — Sign In</Title>
      <TextInput placeholder="Your name" label="Name" value={name} onChange={(e)=>setName(e.target.value)} />
      <Button fullWidth style={{marginTop:16}} onClick={login}>Sign in</Button>
    </Container>
  )
}
