// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "vlq_base128_le.h"

#include <iostream>
#include <fstream>

vlq_base128_le_t::vlq_base128_le_t(kaitai::kstream *p_io, kaitai::kstruct* p_parent, vlq_base128_le_t *p_root) : kaitai::kstruct(p_io) {
    m__parent = p_parent;
    m__root = this;
    f_len = false;
    f_value = false;
    m_groups = new std::vector<group_t*>();
    {
        group_t* _;
        do {
            _ = new group_t(m__io, this, m__root);
            m_groups->push_back(_);
        } while (!(!_->has_next()));
    }
}

vlq_base128_le_t::~vlq_base128_le_t() {
    for (std::vector<group_t*>::iterator it = m_groups->begin(); it != m_groups->end(); ++it) {
        delete *it;
    }
    delete m_groups;
}

vlq_base128_le_t::group_t::group_t(kaitai::kstream *p_io, vlq_base128_le_t* p_parent, vlq_base128_le_t *p_root) : kaitai::kstruct(p_io) {
    m__parent = p_parent;
    m__root = p_root;
    f_has_next = false;
    f_value = false;
    m_b = m__io->read_u1();
}

vlq_base128_le_t::group_t::~group_t() {
}

bool vlq_base128_le_t::group_t::has_next() {
    if (f_has_next)
        return m_has_next;
    m_has_next = (b() & 128) != 0;
    f_has_next = true;
    return m_has_next;
}

int32_t vlq_base128_le_t::group_t::value() {
    if (f_value)
        return m_value;
    m_value = (b() & 127);
    f_value = true;
    return m_value;
}

int32_t vlq_base128_le_t::len() {
    if (f_len)
        return m_len;
    m_len = groups()->size();
    f_len = true;
    return m_len;
}

int32_t vlq_base128_le_t::value() {
    if (f_value)
        return m_value;
    m_value = (((((((groups()->at(0)->value() + ((len() >= 2) ? ((groups()->at(1)->value() << 7)) : (0))) + ((len() >= 3) ? ((groups()->at(2)->value() << 14)) : (0))) + ((len() >= 4) ? ((groups()->at(3)->value() << 21)) : (0))) + ((len() >= 5) ? ((groups()->at(4)->value() << 28)) : (0))) + ((len() >= 6) ? ((groups()->at(5)->value() << 35)) : (0))) + ((len() >= 7) ? ((groups()->at(6)->value() << 42)) : (0))) + ((len() >= 8) ? ((groups()->at(7)->value() << 49)) : (0)));
    f_value = true;
    return m_value;
}
