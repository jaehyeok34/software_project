# software_project

## 사용한 패턴
1. Decorator Pattern
    - system을 확장하는 방법에서 사용함
    - ex: ChatSystem의 단순 echoing 기능을 확장하여 현재 시간도 함께 출력하는 TimeStampDecorator를 생성하여 확장함

2. Factory Pattern
    - System을 생성하는 과정에서 new 키워드 대신 생성할 수 있게 함
    - ex: ChatSystem을 생성하는 Factory 함수로 NewChatSystem()을 구현해서 ChatSystem을 생성함

3. Command Pattern
    - Room이 소유한 System의 호출을 Process() 함수로 정의하고, 인자로 SystemType을 전달 받아 행동을 변경함
